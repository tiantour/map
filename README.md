# map
qq map package for golang

### ToLocation

    package main

    import (
        "fmt"

        "github.com/tiantour/map/qq"
    )

    func main() {
        qq.Key = "xxxx-yyyy-zzzz"
        name := "厦门金宝大酒店"
        location, err := qq.NewQQ().ToLocation(name)
        fmt.Println(location, err)
    }

### ToAddress

    package main

    import (
        "fmt"

        "github.com/tiantour/map/qq"
    )

    func main() {
        qq.Key = "xxxx-yyyy-zzzz"
        lng := 123.45
        lat := 23.45
        address, err := qq.NewQQ().ToAddress(lng, lat)
        fmt.Println(address, err)
    }

### Distance

    package main

    import (
        "fmt"

        "github.com/tiantour/map/earth"
    )

    func main() {
        lng1 := 123.45
        lat1 := 23.45
        lng2 := 122.22
        lat2 := 22.22
        distance := earth.NewEarth().Distance(lat1, lng1, lat2, lng2)
        fmt.Println(distance)
    }