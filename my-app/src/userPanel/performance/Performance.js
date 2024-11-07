import { Component} from "react";

import './Performance.css';
import logo from '../../logo.svg';
import Request from "../../api-backend/request";
import left_arrow from '../../backleftarrowcircularbutton_79585.png';
import right_arrow from '../../rightarrowcircularbutton_79554 (1).png';
import Auth from "../auth/Auth";
import { useParams, Link } from "react-router-dom";



function GetPerformance() {
    const { id_theatre } = useParams();
    return (
        <Performance id={id_theatre} />
    );
}


class Performance extends Component{
    constructor(){
        super();
        this.getDataPerformance = this.getDataPerformance.bind(this);
        this.prepareMap = this.prepareMap.bind(this);
        this.forwardPage = this.forwardPage.bind(this);
        this.backPage = this.backPage.bind(this);
        this.changeLocality = this.changeLocality.bind(this);
        this.updateSizeTheatres = this.updateSizeTheatres.bind(this);
        this.changeDate = this.changeDate.bind(this);
        this.prepareMapPerformance = this.prepareMapPerformance.bind(this);
    }
    
    state = {
        map_performance: [],
        map_performance_formed: [],
        map_type_performance: [],
        map_list_performance: [],
        count_begin: 0,
        count_end: 3,
        size_theatres: 4,
        current_type_performance: 1,
        id_theatre: 0,
        current_date: new Date().toISOString().slice(0,10)
    }

    componentDidMount(){
        this.getDataPerformance();
    }

    render(){
        return <div>
        <div className="mainPerformance">
            <select onChange={e => this.changeLocality(e)} onClick={this.updateSizeTheatres} >
                {
                    this.state.map_type_performance.map(item => (
                        <option key={item.Id_type_performance}>{item.Name}</option>
                    ))
                }
                
            </select>

            <input type="date" onChange={e => this.changeDate(e)} ></input>
            <Link to = {`/`} >
                            <button > 
                                Назад
                            </button>
            </Link>
        </div>

        <div className="mainPerformance">

                <button onClick={this.backPage}>
                    <img src={left_arrow}  alt="left_arrow" />
                </button>
                {
                    this.state.map_performance_formed.map((item, key) => (
                       (this.state.count_begin <= key && this.state.count_end > key) &&
                        <Link to = {`/hall/${this.props.id}`} key = {item.Id_performance}>
                            <button style={{backgroundImage: `url(img/performance/${item.Id_performance}.jpg)`}} className="divPerformance" key = {item.Id_performance} > 
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
            size_theatres: this.state.map_performance_formed.length
        })
    }


    changeDate(e){
        this.setState({
            current_date: e.target.value,
            map_performance_formed: this.prepareMapPerformance(this.state.map_performance, this.state.current_type_performance, 
                this.state.map_list_performance, this.state.current_date, this.state.id_theatre),
        }, () => this.setState({
            current_date: e.target.value,
            map_performance_formed: this.prepareMapPerformance(this.state.map_performance, this.state.current_type_performance, 
                this.state.map_list_performance, this.state.current_date, this.state.id_theatre),
        }))
        
    }

    /**
    Change locality in combobox localities
    */
    changeLocality(e){
        let new_id_type_performance;
        this.state.map_type_performance.forEach((value, key) => {
            if (value.Name== e.target.value){
                new_id_type_performance = value.Id_type_performance;
            }
        })

        this.setState({
            current_type_performance: new_id_type_performance,
            count_begin: 0,
            count_end: 3,
            map_performance_formed: this.prepareMapPerformance(this.state.map_performance, this.state.current_type_performance, 
                this.state.map_list_performance, this.state.current_date, this.state.id_theatre),
        }, () => this.setState({
            current_type_performance: new_id_type_performance,
            count_begin: 0,
            count_end: 3,
            map_performance_formed: this.prepareMapPerformance(this.state.map_performance, this.state.current_type_performance, 
                this.state.map_list_performance, this.state.current_date, this.state.id_theatre),
        }))
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
    Prepare array arays performances
    */
    prepareArrayTheatre(json_request){
        let array = new Array();
        json_request.forEach((i, item) => (
            i.forEach( j => (
                array.push([j['Id_performance'], j['Name'], j['Id_type_performance']])
            ))
        ))
        return array;
    }

    /** 
     * Prepare array maps performances
     */
    prepareMap(json_request){
        let map = new Map(json_request.at(0))
        let array_map = new Array();
        map.forEach((val) => {
            array_map.push(val);
        })
        return array_map;
    }

    /** 
     * Prepare array maps performances by type performance
     */
    prepareMapPerformance(map_performance, current_type_performance, map_list_performance, 
        current_date, id_theatre){
        let map_formed = [];
        let list_performance = [];
        map_list_performance.forEach((val) => {
            if(val.Id_theatre == id_theatre && val.Date == current_date){
                list_performance.push(val.Id_performance);
                //console.log(current_date, "  -----  ", val.Date)
                //console.log(list_performance)

            }
                
        })

  

        map_performance.forEach((val) => {
            //console.log(val.Id_type_performance, "-------------", id_type_performance)
            if(val.Id_type_performance == current_type_performance ){
                list_performance.forEach((id) => {
                    if(id == val.Id_performance){
                        map_formed.push(val);
                        //console.log(val);
                    }
                })

            }
                
        })
        //console.log(list_performance, " list_performance");
        //console.log(map_formed, " map_performance");
        return map_formed;
    }


    /** 
     * Get type performance and performances from database.
     */
    async getDataPerformance(){

        let request = new Request().post("/theatre/performance");
        let json_performance = await Promise.all([request]);

        let request2 = new Request().post("/theatre/performance/type_performance");
        let json_type_performance = await Promise.all([request2]);

        let request3 = new Request().post("/theatre/performance/list_performance");
        let json_list_performance = await Promise.all([request3]);
        

        let id_theatre = this.props.id;
        let map_performance = this.prepareMap(json_performance);
        let map_list_performance = this.prepareMap(json_list_performance);
        let map_type_performance = this.prepareMap(json_type_performance);
        let map_performance_formed = this.prepareMapPerformance(map_performance, this.state.current_type_performance, 
            map_list_performance, this.state.current_date, id_theatre);

        this.setState({
            id_theatre: id_theatre,
            map_performance: map_performance,
            map_list_performance: map_list_performance,
            map_type_performance: map_type_performance,
            map_performance_formed: map_performance_formed
        })
     }
}


export default GetPerformance;