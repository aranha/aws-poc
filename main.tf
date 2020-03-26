provider "aws" {
    region = "us-east-1"
    access_key = "${var.aws_access_key}"
    secret_key = "${var.aws_secret_key}"
}

data "aws_ami" "aws_ami_centos" {
    owners = ["self"]
    filter {
        name = "name"
        values = ["ami tema 13 devops"]
    }
}

resource "aws_instance" "devops-instance" {
    key_name = "kp_tema13_devops"
    ami = "${data.aws_ami.aws_ami_centos.id}"
    instance_type = "t2.micro"
    tags = {
        Name = "tema13-devops"
    }
}

resource "aws_security_group" "sg-devops" {
    name = "security group tema 13 devops"
    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["191.32.54.163/32"]
    }
    ingress {
        from_port = 5000
        to_port = 5000
        protocol = "tcp"
        cidr_blocks = ["191.32.54.163/32"]
    }
}

resource "aws_launch_configuration" "lc-devops" {
    name_prefix = "lc-tema13-devops"
    image_id = "${data.aws_ami.aws_ami_centos.id}"
    instance_type = "t2.micro"
    key_name = "kp_tema13_devops"
    security_groups = ["${aws_security_group.sg-devops.id}"]
}

resource "aws_autoscaling_group" "asg-devops" {
    name = "asg-devops-tema13"
    launch_configuration = "${aws_launch_configuration.lc-devops.name}"
    max_size = 1
    min_size = 1
    availability_zones = ["us-east-1a"]
    load_balancers = ["${aws_elb.elb-devops.name}"]
    tag {
        key = "Name"
        value = "devops-tema13"
        propagate_at_launch = true
    }
}

resource "aws_elb" "elb-devops" {
    name = "elb-tema13-devops"
    availability_zones = ["us-east-1a"]
    listener {
        instance_port = 5000
        instance_protocol = "http"
        lb_port = 5000
        lb_protocol = "http"
    }
    instances = "${aws_instance.devops-instance.*.id}"
    source_security_group = "${aws_security_group.sg-devops.id}"
}