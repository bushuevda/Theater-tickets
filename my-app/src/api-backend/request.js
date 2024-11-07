/*     
    fetch("http://127.0.0.1:8000/", {headers:{"Origin": "http://localhost:8000"}})
    .then(response => response.text())
    .then(data => console.log(data))
    .catch(error => console.error(error));
 */

import formed_json from "./formed_json"

class Request{
    constructor(){
        this.origin_url_ = "http://localhost:8000";
        this.post = this.post.bind(this);
        this.get = this.get.bind(this);
    }   

     async post(url_to){
        let request = await fetch(this.origin_url_ + url_to, {headers:{"Origin": this.origin_url_}});
        let jsonData = await request.json();
        jsonData = await formed_json(jsonData);
        return jsonData
    }
    async get(url_to){
        let request = await fetch(this.origin_url_ + url_to, {headers:{"Origin": this.origin_url_}});
        let jsonData = await request.json();
        return await formed_json(jsonData);
    }
}

export default Request;