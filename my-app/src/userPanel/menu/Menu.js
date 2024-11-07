import { Component } from "react";
import './Menu.css';
import Auth from "../auth/Auth"
import Theatre from "../theatre/Theatre";
import Performance from "../performance/Performance"

class Menu extends Component{
    constructor(){
        super();
        this.btnAuth = this.btnAuth.bind(this);
    }
    state = {
        auth: {state_press: false, component: ""},
        theatre: {state_press: false, component: <Theatre/>},
        performance: {state_press: false, component: <Performance/>},

    }
    render(){
        return <div className="mainBtns">
            <button className="btnAuth menuBtn" onClick={this.btnAuth}>Авторизация</button>
            {this.state.auth.component}
            {this.state.theatre.component}
        </div>
    }

    btnAuth(){
        if (this.state.auth.state_press)
            this.setState({
                auth: {state_press: false, component: ""}
            });
        else
            this.setState({
                auth: {state_press: true, component: <Auth/>}
            });
        console.log('s')
    }
    btnTheatre(){

    }
}

export default Menu;