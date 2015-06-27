// +build integration

package main

import (
    "github.com/sethgrid/go-supertest"
    "testing"
)

func wrapper() {
    go main()
}

func TestGetContent(t *testing.T) {

    app := supertest.NewAppRunner("localhost:3000", wrapper)

    // start your app
    app.Start()

    // you can method chain and check headers, content, and status
    err := app.Get("/v1/api/version").
        ExpectStatusCode("200 OK").
        ExpectHeader("Content-Type", "text/plain; charset=utf-8").
        ExpectContent([]byte("1.0")).
        End()
    if err != nil {
        t.Error("Error Getting Version", err)
    }
}
