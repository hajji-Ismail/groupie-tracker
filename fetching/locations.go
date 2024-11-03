package fetching

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"

    "groupie-tracker/models"
)

var local []models.Location

func init() {
    var err error
    local, err = fetchlocal("https://groupietrackers.herokuapp.com/api/locations")
    if err != nil {
        log.Fatal(err)
    }
}

func fetchlocal(url string) ([]models.Location, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch data: status code %d", response.StatusCode)
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var loc []models.Location
    err = json.Unmarshal(body, &loc)
    if err != nil {
        return nil, err
    }
    return loc, nil
}