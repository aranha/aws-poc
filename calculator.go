package main

import (
        "encoding/json"
        "fmt"
        "github.com/gorilla/mux"
        "net/http"
        "strconv"
)

var historic []*Operation

type Operation struct {
        Number1 float64
        Signal  string
        Number2 float64
        Result  float64
}

func sum(num1 float64, num2 float64) float64 {
        return num1 + num2
}

func sub(num1 float64, num2 float64) float64 {
        return num1 - num2
}

func mult(num1 float64, num2 float64) float64 {
        return num1 * num2
}

func div(num1 float64, num2 float64) float64 {
        return num1 / num2
}

func calc(response http.ResponseWriter, request *http.Request){
        operation := new(Operation)
        params1 := mux.Vars(request)["operation"]

        num1, err := strconv.ParseFloat(mux.Vars(request)["number1"], 64)
        num2, err1 := strconv.ParseFloat(mux.Vars(request)["number2"], 64)
        
        if err != nil {
                response.WriteHeader(http.StatusExpectationFailed)
                fmt.Fprint(response, "num1 is not a number")
                return
        }

        if err1 != nil {
                response.WriteHeader(http.StatusExpectationFailed)
                fmt.Fprint(response, "num2 is not a number")
                return
        }

        if params1 == "div" && num2 == 0 {
                response.WriteHeader(http.StatusExpectationFailed)
                fmt.Fprint(response, "cannot divide by zero")
                return
        }

        var result float64
        switch params1 {
        case "sum":
                params1 = "+"
                result = sum(num1, num2)
        case "sub":
                params1 = "-"
                result = sub(num1, num2)
        case "mult":
                params1 = "*"
                result = mult(num1, num2)
        case "div":
                params1 = "/"
                result = div(num1, num2)
        default:
              response.WriteHeader(http.StatusExpectationFailed)
              fmt.Fprint(response, "operation's not available or doesn't exist")
              return
        }

        operation.Number1 = num1
        operation.Number2 = num2
        operation.Signal = params1
        operation.Result = result

        operationJson, err := json.Marshal(operation)
        historic = append(historic, operation)

        response.Write(operationJson)

}

func hist(response http.ResponseWriter, request *http.Request){
        response.Header().Set("Content-Type", "application/json")
        jsonHistory, err := json.Marshal(historic)
        if err != nil{
                fmt.Fprint(response, err)
        }
        response.Write(jsonHistory)
}

func main() {
        fmt.Print("Serving at http://0.0.0.0:5000\n")
        routes := mux.NewRouter()
        calcApi := routes.PathPrefix("/calc").Subrouter()
        calcApi.HandleFunc("/{operation}/{number1}/{number2}", calc)
        calcApi.HandleFunc("/hist", hist)
        http.ListenAndServe(":5000", routes)
}