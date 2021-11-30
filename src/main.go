package main

func main() {
    loadEnv()
    setupDB()
    transmission := getTransmissionClient()
    torrents := getTorrents(transmission)
    processTorrents(torrents)
}
