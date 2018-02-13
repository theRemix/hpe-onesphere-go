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

package osbinding

import (
    //"fmt"
    //"strings"
    "bytes"
    //"io"
    "io/ioutil"
    "net/http"
    //"net/url"
    "encoding/json"
)

var Token string
var HostUrl string

func Connect(hostUrl, user, password string) {
    HostUrl = hostUrl
    fullUrl := hostUrl + "/rest/session"
    values := map[string]string{"userName": user, "password": password}
    jsonValue, err := json.Marshal(values)
    req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonValue))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    //bodyStr := string(body)
    var dat map[string]string
    err = json.Unmarshal(body, &dat)
    if err != nil {
        panic(err)
    }

    Token = dat["token"]
}

func callHttpRequest(method, url string, values interface{}) string {
    jsonValue, err := json.Marshal(values)
    if err != nil {
        panic(err)
    }
    req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", Token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    bodyStr := string(bodyBytes)
    return bodyStr
}

func GetStatus() string {
    fullUrl := HostUrl + "/rest/status"
    return callHttpRequest("GET", fullUrl, nil)
}

// os="windows" or os="mac"
func GetConnectApp(os string) string {
    fullUrl := HostUrl + "/rest/connect-app"
    values := map[string]string{"os": os}
    return callHttpRequest("GET", fullUrl, values)
}

// Session APIs

// view="full"
func GetSession(view string) string {
    fullUrl := HostUrl + "/rest/session"
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func GetSessionIdp(username string) string {
    fullUrl := HostUrl + "/rest/session/idp"
    values := map[string]string{"userName": username}
    return callHttpRequest("GET", fullUrl, values)
}

// Account APIs

// view="full"
func GetAccount(view string) string {
    fullUrl := HostUrl + "/rest/account"
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

// Appliances APIs

func GetAppliances() string {
    fullUrl := HostUrl + "/rest/appliances"
    return callHttpRequest("GET", fullUrl, nil)
}

func CreateAppliance(epAddress, epUsername, epPassword, 
                     name, regionUri, applianceType string) string {
    fullUrl := HostUrl + "/rest/appliances"
    values := map[string]interface{}{
        "endpoint": map[string]interface{}{
            "address": epAddress,
            "password": epPassword,
            "username": epUsername},
        "name": name,
        "regionUri": regionUri,
        "type": applianceType}
    return callHttpRequest("POST", fullUrl, values)
}

func GetAppliance(applianceID string) string {
    fullUrl := HostUrl + "/rest/appliances/" + applianceID
    return callHttpRequest("GET", fullUrl, nil)
}

func DeleteAppliance(applianceID string) string {
    fullUrl := HostUrl + "/rest/appliances/" + applianceID
    return callHttpRequest("DELETE", fullUrl, nil)
}

// infoArray: [{os, path, value}]
func UpdateAppliance(applianceID string, infoArray []string) string {
    fullUrl := HostUrl + "/rest/appliances/" + applianceID
    values := infoArray 
    return callHttpRequest("PUT", fullUrl, values)
}

// Providers APIs

func GetProviderTypes() string {
    fullUrl := HostUrl + "/rest/provider-types"
    return callHttpRequest("GET", fullUrl, nil)
}

func GetProviders(parentUri, providerTypeUri string) string {
    fullUrl := HostUrl + "/rest/providers"
    values := map[string]string{"parentUri": parentUri, "providerTypeUri": providerTypeUri}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateProvider(providerID, providerTypeUri, accessKey, secretKey, 
                    s3CostBucket, parentUri, state string, 
                    paymentProvider bool) string {
    fullUrl := HostUrl + "/rest/providers"
    values := map[string]interface{}{
        "id": providerID,
        "providerTypeUri": providerTypeUri,
        "accessKey": accessKey,
        "secretKey": secretKey,
        "paymentProvider": paymentProvider,
        "s3CostBucket": s3CostBucket,
        "parentUri": parentUri,
        "state": state}
    return callHttpRequest("POST", fullUrl, values)
}

// view="full"
func GetProvider(providerID, view string, discover bool) string {
    fullUrl := HostUrl + "/rest/providers/" + providerID
    values := map[string]interface{}{
        "view": view,
        "discover": discover}
    return callHttpRequest("GET", fullUrl, values)
}

func DeleteProvider(providerID string) string {
    fullUrl := HostUrl + "/rest/providers/" + providerID
    return callHttpRequest("DELETE", fullUrl, nil)
}

func UpdateProvider(providerID, infoArray string) string {
    fullUrl := HostUrl + "/rest/providers/" + providerID
    return callHttpRequest("PUT", fullUrl, infoArray)
}

// Regions APIs

func GetRegions(providerUri, view string) string {
    fullUrl := HostUrl + "/rest/regions"
    values := map[string]string{"providerUri": providerUri, "view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateRegion(name, providerUri, locLatitude, locLongitude string) string {
    fullUrl := HostUrl + "/rest/regions"
    values := map[string]interface{}{
        "location": map[string]interface{}{
            "latitude": locLatitude,
            "longitude": locLongitude},
        "name": name,
        "providerUri": providerUri}
    return callHttpRequest("POST", fullUrl, values)
}

func GetRegion(regionID, view string, discover bool) string {
    fullUrl := HostUrl + "/rest/regions/" + regionID
    values := map[string]interface{}{"view": view, "discover": discover}
    return callHttpRequest("GET", fullUrl, values)
}

func UpdateRegion(regionID, info string) string {
    fullUrl := HostUrl + "/rest/regions/" + regionID
    return callHttpRequest("PUT", fullUrl, info)
}

// Zone Types APIs

func GetZoneTypes() string {
    fullUrl := HostUrl + "/rest/zone-types"
    return callHttpRequest("GET", fullUrl, nil)
}

// Zones APIs

func GetZones(regionUri, query string) string {
    fullUrl := HostUrl + "/rest/zones"
    values := map[string]string{"regionUri": regionUri, "query": query}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateZone(name, providerUri, regionUri, zoneTypeUri string) string {
    fullUrl := HostUrl + "/rest/zones"
    values := map[string]string{"name": name, "providerUri": providerUri, 
                                "regionUri": regionUri, "zoneTypeUri": zoneTypeUri}
    return callHttpRequest("POST", fullUrl, values)
}

func GetZone(zoneID, view string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func UpdateZone(zoneID, infoArray string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID
    return callHttpRequest("PUT", fullUrl, infoArray)
}

func DeleteZone(zoneID string, force bool) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID
    values := map[string]bool{"force": force}
    return callHttpRequest("DELETE", fullUrl, values)
}

func ActionOnZone(zoneID, action string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID + "/actions"
    values := map[string]string{"type": action}
    return callHttpRequest("POST", fullUrl, values)
}

func GetZoneApplianceImage(zoneID string) string {
    fullUrl := HostUrl + "/rest/zones/" + zoneID + "/appliance-image"
    return callHttpRequest("GET", fullUrl, nil)
}

// Catalogs APIs

func GetCatalogs(query string) string {
    fullUrl := HostUrl + "/rest/catalogs"
    values := map[string]string{"q": query}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateCatalog(name, url string) string {
    fullUrl := HostUrl + "/rest/catalogs"
    values := map[string]string{"name": name, "url": url}
    return callHttpRequest("POST", fullUrl, values)
}

func GetCatalog(catalogID string) string {
    fullUrl := HostUrl + "/rest/catalogs/" + catalogID
    return callHttpRequest("GET", fullUrl, nil)
}

func UpdateCatalog(catalogID, name, status, uri, url, serviceTypeUri, 
                   timeCreated, timeModified string) string {
    fullUrl := HostUrl + "/rest/catalogs/" + catalogID
    values := map[string]interface{}{
        "created": timeCreated,
        "id": catalogID,
        "modified": timeModified,
        "name": name,
        "status": status,
        "uri": uri,
        "url": url,
        "serviceTypeUri": serviceTypeUri}
    return callHttpRequest("PUT", fullUrl, values)
}

// Service Types APIs

func GetServiceTypes() string {
    fullUrl := HostUrl + "/service-types"
    return callHttpRequest("GET", fullUrl, nil)
}

func GetServiceType(serviceTypeID string) string {
    fullUrl := HostUrl + "/service-types/" + serviceTypeID
    return callHttpRequest("GET", fullUrl, nil)
}

// Services APIs

func GetServices(query, userQuery, serviceTypeUri, zoneUri, 
                 workspaceUri, catalogUri, view string) string {
    fullUrl := HostUrl + "/rest/services"
    values := map[string]string{
        "query": query, 
        "userQuery": userQuery, 
        "serviceTypeUri": serviceTypeUri, 
        "zoneUri": zoneUri, 
        "workspaceUri": workspaceUri, 
        "catalogUri": catalogUri, 
        "view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func GetService(serviceID, view string) string {
    fullUrl := HostUrl + "/rest/services/" + serviceID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

// Virtual Machine Profiles APIs

func GetVirtualMachineProfiles(query, zoneUri, serviceUri string) string {
    fullUrl := HostUrl + "/rest/virtual-machine-profiles"
    values := map[string]string{"q": query, "zoneUri": zoneUri, "serviceUri": serviceUri}
    return callHttpRequest("GET", fullUrl, values)
}

func GetVirtualMachineProfile(vmProfileID string) string {
    fullUrl := HostUrl + "/rest/virtual-machine-profiles/" + vmProfileID
    return callHttpRequest("GET", fullUrl, nil)
}

// Networks APIs

func GetNetworks(query, zoneUri string) string {
    fullUrl := HostUrl + "/rest/networks"
    values := map[string]string{"q": query, "zoneUri": zoneUri}
    return callHttpRequest("GET", fullUrl, values)
}

func GetNetwork(networkID string) string {
    fullUrl := HostUrl + "/rest/networks/" + networkID
    return callHttpRequest("GET", fullUrl, nil)
}

// Workspaces APIs

func GetWorkspaces(query, view string) string {
    fullUrl := HostUrl + "/rest/workspaces"
    values := map[string]string{"q": query, "view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateWorkspace(name, description, tagUrisArray string) string {
    fullUrl := HostUrl + "/rest/workspaces"
    values := map[string]string{"name": name, "description": description, "tagUris": tagUrisArray}
    return callHttpRequest("POST", fullUrl, values)
}

func GetWorkspace(workspaceID, view string) string {
    fullUrl := HostUrl + "/rest/workspaces/" + workspaceID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func UpdateWorkspace(workspaceID, name, description, tagUrisArray string) string {
    fullUrl := HostUrl + "/rest/workspaces/" + workspaceID
    values := map[string]string{"name": name, "description": description, "tagUris": tagUrisArray}
    return callHttpRequest("PUT", fullUrl, values)
}

func DeleteWorkspace(workspaceID string) string {
    fullUrl := HostUrl + "/rest/workspaces/" + workspaceID
    return callHttpRequest("DELETE", fullUrl, nil)
}

// Deployments APIs

func GetDeployments(query, view string) string {
    fullUrl := HostUrl + "/rest/deployments"
    values := map[string]string{"query": query, "view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateDeployment(info string) string {
    fullUrl := HostUrl + "/rest/deployments"
    return callHttpRequest("POST", fullUrl, info)
}

func GetDeployment(deploymentID, view string) string {
    fullUrl := HostUrl + "/rest/deployments/" + deploymentID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func UpdateDeployment(deploymentID, info string) string {
    fullUrl := HostUrl + "/rest/deployments/" + deploymentID
    return callHttpRequest("PUT", fullUrl, info)
}

func DeleteDeployment(deploymentID string) string {
    fullUrl := HostUrl + "/rest/deployments/" + deploymentID
    return callHttpRequest("DELETE", fullUrl, nil)
}

func ActionOnDeployment(deploymentID, actionType, force string) string {
    fullUrl := HostUrl + "/rest/deployments/" + deploymentID + "/actions"
    values := map[string]string{"force": string(force), "type": actionType}
    return callHttpRequest("POST", fullUrl, values)
}

func GetDeploymentConsole(deploymentID string) string {
    fullUrl := HostUrl + "/rest/deployments/" + deploymentID + "/console"
    return callHttpRequest("GET", fullUrl, nil)
}

// Memberships APIs

func GetMemberships(query string) string {
    fullUrl := HostUrl + "/rest/memberships"
    values := map[string]string{"query": query}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateMembership(userUri, roleUri, workspaceUri string) string {
    fullUrl := HostUrl + "/rest/memberships"
    values := map[string]string{"userUri": userUri, "roleUri": roleUri, "workspaceUri": workspaceUri}
    return callHttpRequest("POST", fullUrl, values)
}

func DeleteMembership(userUri, roleUri, workspaceUri string) string {
    fullUrl := HostUrl + "/rest/memberships"
    values := map[string]string{"userUri": userUri, "roleUri": roleUri, "workspaceUri": workspaceUri}
    return callHttpRequest("DELETE", fullUrl, values)
}

// Roles APIs

func GetRoles() string {
    fullUrl := HostUrl + "/rest/roles"
    return callHttpRequest("GET", fullUrl, nil)
}

// Users APIs

func GetUsers() string {
    fullUrl := HostUrl + "/rest/users"
    return callHttpRequest("GET", fullUrl, nil)
}

func CreateUser(name, password, email, role string) string {
    fullUrl := HostUrl + "/rest/users"
    values := map[string]string{"name": name, "email": email, "password": password, "role": role}
    return callHttpRequest("POST", fullUrl, values)
}

func GetUser(userID string) string {
    fullUrl := HostUrl + "/rest/users/" + userID
    return callHttpRequest("GET", fullUrl, nil)
}

func UpdateUser(userID, name, password, email, role string) string {
    fullUrl := HostUrl + "/rest/users/" + userID
    values := map[string]string{"name": name, "email": email, "password": password, "role": role}
    return callHttpRequest("PUT", fullUrl, values)
}

func DeleteUser(userID string) string {
    fullUrl := HostUrl + "/rest/users/" + userID
    return callHttpRequest("DELETE", fullUrl, nil)
}

// Metrics APIs

func GetMetrics(
        resourceUriArray, categoryArray, queryArray, nameArray []string, 
        periodStart string,
        period, periodCount int,
        view string,
        start string,
        count int) string {
    fullUrl := HostUrl + "/rest/metrics"
    values := map[string]interface{}{
        "resourceUri": resourceUriArray,
        "category": categoryArray,
        "query": queryArray,
        "nameArray": nameArray,
        "periodStart": periodStart,
        "period": period,
        "periodCount": periodCount,
        "view": view,
        "start": start,
        "count": count}
    return callHttpRequest("GET", fullUrl, values)
}

// Events APIs

func GetEvents(resourceUri string) string {
    fullUrl := HostUrl + "/rest/events"
    values := map[string]string{"resourceUri": resourceUri}
    return callHttpRequest("GET", fullUrl, values)
}

// Volumes APIs

func GetVolumes(query string) string {
    fullUrl := HostUrl + "/rest/volumes"
    values := map[string]string{"query": query}
    return callHttpRequest("GET", fullUrl, values)
}

func GetVolume(volumeID string) string {
    fullUrl := HostUrl + "/rest/volumes/" + volumeID
    return callHttpRequest("GET", fullUrl, nil)
}

// Tag Keys APIs

func GetTagKeys(view string) string {
    fullUrl := HostUrl + "/rest/tag-keys"
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateTagKey(name string) string {
    fullUrl := HostUrl + "/rest/tag-keys"
    values := map[string]string{"name": name}
    return callHttpRequest("POST", fullUrl, values)
}

func GetTagKey(tagKeyID, view string) string {
    fullUrl := HostUrl + "/rest/tag-keys/" + tagKeyID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func DeleteTagKey(tagKeyID string) string {
    fullUrl := HostUrl + "/rest/tag-keys/" + tagKeyID
    return callHttpRequest("DELETE", fullUrl, nil)
}

// Tags APIs

func GetTags(view string) string {
    fullUrl := HostUrl + "/rest/tags"
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func CreateTag(name, tagKeyUri string) string {
    fullUrl := HostUrl + "/rest/tags"
    values := map[string]string{"name": name, "tagKeyUri": tagKeyUri}
    return callHttpRequest("POST", fullUrl, values)
}

func GetTag(tagID, view string) string {
    fullUrl := HostUrl + "/rest/tags/" + tagID
    values := map[string]string{"view": view}
    return callHttpRequest("GET", fullUrl, values)
}

func DeleteTag(tagID string) string {
    fullUrl := HostUrl + "/rest/tags/" + tagID
    return callHttpRequest("DELETE", fullUrl, nil)
}

// Keypairs APIs

func GetKeyPair(regionUri, workspaceUri string) string {
    fullUrl := HostUrl + "/rest/keypairs"
    values := map[string]string{"regionUri": regionUri, "workspaceUri": workspaceUri}
    return callHttpRequest("GET", fullUrl, values)
}

