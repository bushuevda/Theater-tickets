import { Component} from "react";

import './Theatre.css' ;
import Request from "../../api-backend/request";
import left_arrow from '../../backleftarrowcircularbutton_79585.png';
import right_arrow from '../../rightarrowcircularbutton_79554 (1).png';
import {Link } from "react-router-dom";


class Theatre extends Component{
    constructor(){
        super();
        this.getDataTheatre = this.getDataTheatre.bind(this);
        this.prepareMapTheatre = this.prepareMapTheatre.bind(this);
        this.forwardPage = this.forwardPage.bind(this);
        this.backPage = this.backPage.bind(this);
        this.changeLocality = this.changeLocality.bind(this);
        this.updateSizeTheatres = this.updateSizeTheatres.bind(this);
    }
    state = {
        map_theatre: [],
        map_theatre_by_loc: [],
        map_locality: [],
        count_begin: 0,
        count_end: 3,
        size_theatres: 4,
        current_locality: 1
    }
    componentDidMount(){
        this.getDataTheatre();
    }
    render(){
        return <div>
        <div className="mainPerformance">
            <select onChange={e => this.changeLocality(e)} onClick={this.updateSizeTheatres} style={{fontSize:'25px'}}>
                {
                    this.state.map_locality.map(item => (
                        <option key={item.Id_locality}>{item.Name}</option>
                    ))
                }
                
            </select>
        </div>

        <div className="mainPerformance" >

                <button onClick={this.backPage} >
                    <img src={left_arrow}  alt="left_arrow" />
                </button>
                {
                    this.state.map_theatre_by_loc.map((item, key) => (
                       (this.state.count_begin <= key && this.state.count_end > key) &&
                        <Link to = {`/theatre/${item.Id_theatre}`} key = {item.Id_theatre}>
                            <button style={{backgroundImage: `url(img/theatres/${item.Id_theatre}.jpg)`}}  className="divPerformance" key = {item.Id_theatre} > 
                            </button>
                            <p style={{color: 'black', fontSize: '34px', fontWeight: 'bold',  textDecoration: 'inherit'}}>{item.Name}</p>
                        </Link>
                    ))

                }            
                <button onClick={this.forwardPage}>
                    <img src={right_arrow}  alt="right_arrow" />
                </button>
        </div>

      </div>
    }


    /**
    Update size array maps theatres by locality
    */
    updateSizeTheatres(){
        this.setState({
            size_theatres: this.state.map_theatre_by_loc.length
        })
    }

    /**
    Change locality in combobox localities
    */
    changeLocality(e){

        let new_id_locality;
        this.state.map_locality.forEach((value, key) => {
            if (value.Name== e.target.value){
                new_id_locality = value.Id_locality;
            }
        })

        this.setState({
            current_locality: new_id_locality,
            count_begin: 0,
            count_end: 3,
            map_theatre_by_loc: this.prepareMapTheatreByLoc(this.state.map_theatre, new_id_locality),
        })
    }

    /**
    Shift theatres forward
    */
    forwardPage(){
        if (this.state.size_theatres > this.state.count_end && this.state.size_theatres > 3)
            this.setState({
                count_begin: this.state.count_begin + 1,
                count_end: this.state.count_end + 1
            })
    }

    /**
    Shift theatres back
    */
    backPage(){
        if (this.state.count_begin > 0 && this.state.size_theatres > 3)
            this.setState({
                count_begin: this.state.count_begin - 1,
                count_end: this.state.count_end - 1
            })
    }

    /**
    Prepare array arays theatres
    */
    prepareArrayTheatre(json_request){
        let array = new Array();
        json_request.forEach((i, item) => (
            i.forEach( j => (
                array.push([j['Id_theatre'], j['Name'], j['Id_locality']])
            ))
        ))
        return array;
    }

    /** 
     * Prepare array maps theatres
     */
    prepareMapTheatre(json_request){
        let map = new Map(json_request.at(0))
        let array_map = new Array();
        map.forEach((val) => {
            array_map.push(val);
        })
        return array_map;
    }

    /** 
     * Prepare array maps theatres by locality
     */
    prepareMapTheatreByLoc(map_theatre, current_locality){
        let map_temp = []
        map_theatre.forEach((val) => {
            if(val.Id_locality == current_locality)
                map_temp.push(val);
        })
        return map_temp;
    }

    /** 
     * Get localities and theatres from database.
     */
    async getDataTheatre(){
        let request = new Request().post("/theatre/");
        let json_theatre = await Promise.all([request]);

        this.setState({
            map_theatre: this.prepareMapTheatre(json_theatre)
        });

        let request2 = new Request().post("/theatre/locality/");
        let json_locality = await Promise.all([request2]);

        this.setState({
            map_locality: this.prepareMapTheatre(json_locality)
        });

        this.setState({
            map_theatre_by_loc: this.prepareMapTheatreByLoc(this.state.map_theatre, this.state.current_locality)
        })

     }
}

export default Theatre;