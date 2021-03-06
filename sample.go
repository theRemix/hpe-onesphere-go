// (C) Copyright 2018 Hewlett Packard Enterprise Development LP.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
    osbinding "./onesphere"
    "fmt"
)

func main() {
    osbinding.Connect("https://onesphere-host-url", "username", "password")
    fmt.Println("Token:", osbinding.Token)

    fmt.Println("Status:", osbinding.GetStatus())
    fmt.Println("Session:", osbinding.GetSession("full"))
    fmt.Println("Account:", osbinding.GetAccount("full"))
    fmt.Println("ProviderTypes:", osbinding.GetProviderTypes())
    fmt.Println("ZoneTypes:", osbinding.GetZoneTypes())
    fmt.Println("ServiceTypes:", osbinding.GetServiceTypes())
    fmt.Println("Roles:", osbinding.GetRoles())
    fmt.Println("Users:", osbinding.GetUsers())
    fmt.Println("TagKeys:", osbinding.GetTagKeys("full"))
    fmt.Println("Tags:", osbinding.GetTags("full"))

    osbinding.Disconnect()
}

