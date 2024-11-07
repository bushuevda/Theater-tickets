import './Auth.css';

import { Component } from 'react';

class Auth extends Component{
    render(){
        return <div className="divAuth">
            <p>Логин</p>
            <input type="text"></input><br></br>
            <p>Пароль</p>
            <input type="password"></input><br></br>
            <input type="submit"></input>
        </div>
    }
}

export default Auth;