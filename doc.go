// The MIT License (MIT)

// Copyright (c) 2016 Claudemiro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

/*
Package health is a easy to use, extensible health check library.

        package main

        import (
            "net/http"

            "github.com/dimiro1/health"
            "github.com/dimiro1/health/url"
        )

        func main() {
            companies := health.NewCompositeChecker()
            companies.AddChecker("Microsoft", url.NewChecker("https://www.microsoft.com/"))
	        companies.AddChecker("Oracle", url.NewChecker("https://www.oracle.com/"))
	        companies.AddChecker("Google", url.NewChecker("https://www.google.com/"))

            handler := health.NewHandler()
            handler.AddChecker("Go", url.NewChecker("https://golang.org/"))
            handler.AddChecker("Big Companies", companies)

            http.Handle("/health/", handler)
            http.ListenAndServe(":8080", nil)
        }
*/
package health
