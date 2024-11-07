import './App.css';
import { Component } from 'react';
import {BrowserRouter, Route, Routes} from 'react-router-dom';
import Menu from './userPanel/menu/Menu'
import Theatre from './userPanel/theatre/Theatre';
import GetPerformance from './userPanel/performance/Performance';
import GetHall from './userPanel/hall/Hall';


class App extends Component{
  render (){
    return (<div className='App'>
    <header className='App-header'>
      <BrowserRouter>
          <div>
              <Routes>
                  <Route path="/" element={<Theatre />} />
                  <Route path="/theatre/:id_theatre" Component = {GetPerformance} />
                  <Route path="/hall/:id_theatre" Component={GetHall}/>
              </Routes>
          </div>
      </BrowserRouter>
      </header>
      </div>
  );
  };
}

export default App;
