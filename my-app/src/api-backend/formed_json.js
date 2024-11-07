function formed_json(jsonData){
    jsonData = JSON.stringify(jsonData);
    jsonData = JSON.parse(jsonData);
    jsonData = new Map(Object.entries(jsonData));
    return jsonData;
}

export default formed_json;