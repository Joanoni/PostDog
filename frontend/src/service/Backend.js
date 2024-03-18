import axios from "axios";

export default class Backend {

    static scheme = "http://"
    static server = "localhost"
    static port = ":7963/"

    static url() {
        return Backend.scheme + Backend.server + Backend.port
    }

    static request(method, url, payload) {
        console.log("payload:", payload)
        return axios({
            method: method,
            url: Backend.url() + url,
            data: payload
        })
    }
}
