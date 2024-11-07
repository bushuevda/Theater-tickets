import {Component} from "react";
import { useParams, Link } from "react-router-dom";
import Request from "../../api-backend/request";

function GetHall() {
    const { id_theatre } = useParams();
    return (
        <Hall id={id_theatre} />
    );
}



class Hall extends Component{
    constructor(){
        super();
        this.getDataHall = this.getDataHall.bind(this);
        this.setNumberPlace =this.setNumberPlace.bind(this);
    }

    state = {
        map_hall_places: [],
        number_place: 0
    }
    render(){
        return <div className="divAuth">
            <div>
           
                {this.state.map_hall_places.map((item, key) => (
                    <div key={key}>
                        {item.Places.map((it, k) => (
                            <button key={k} value={`${it.Number}, ${it.Id_row}, ${it.Id_place}` } onClick={(e) => this.setNumberPlace(e)}>
                                {it.Number}
                            </button>
                        ))}
                    </div>
                    
                ))
                }
                  
       
            <div>
                <input type="text" value={this.state.number_place} onChange={this.changeInpText}></input>
            </div>
 
            </div>
        </div>
    }
    componentDidMount(){
        console.log(this.props.id);
        this.getDataHall();
    }

    componentDidCatch(){
        
    }


    changeInpText(){

    }

    setNumberPlace(e){
        console.log(e.target.value)
        this.setState({
            number_place: e.target.value,
        },
        )
    }

    /** 
     * Prepare array maps 
     */
    prepareMap(json_request){
        let map = new Map(json_request.at(0))
        let array_map = new Array();
        map.forEach((val) => {
            array_map.push(val);
        })
        return array_map;
    }

    async prepareMapPlaces(prepared_json_rows){
        let map = []//[{"1":{"Id_row":1,"Number":1,"Id_hall":1, "Places":{"1":{}, "2":{}}}}];

        for(const item of prepared_json_rows){
            try {
                let pl = new Request().post(`/theatre/performance/hall/rows/places/${item.Id_row}`);
                let res = await Promise.all([pl]);
                let array_places = []
                res.forEach(val => {
                    val.forEach(it => {
                        array_places.push(it)
                    })
                })
                let m = {"Id_row": item.Id_row, "Number":item.Number, "Id_hall":item.Id_hall, "Places":array_places};
                map.push(m);
            } catch (error) {
            }
        }
        return map.at(0);
    }

    async getDataHall(){
        let request = new Request().post(`/theatre/performance/hall/${this.props.id}`);
        let json_hall = await Promise.all([request]);

        let id_hall = this.prepareMap(json_hall)[0].Id_hall;
        //console.log(id_hall, " --- id_hall");

        let request2 = new Request().post(`/theatre/performance/hall/rows/${id_hall}`);
        let json_rows = await Promise.all([request2]);
        let prepared_json_rows = this.prepareMap(json_rows);


        let prepared_places = await this.prepareMapPlaces(prepared_json_rows)
        let map_hall_places = await Promise.all([prepared_places]);
        this.setState({
            map_hall_places: map_hall_places
        })
        console.log(this.state.map_hall_places)
    }
}

export default GetHall;