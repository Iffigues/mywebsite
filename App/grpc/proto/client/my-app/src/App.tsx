import React, { useState, useRef, useLayoutEffect}  from 'react';
import logo from './logo.svg';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import { Navbar, Container, Nav, Form, FormControl,Button } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Image from 'react-bootstrap/Image'
import Dropdown from 'react-bootstrap/Dropdown';
import './App.css';

import Navi from './Navi'
import Fortune from './Fortune'

function Hellso() {
  return (
    <p>Hello</p>
  );
}

const Hello = React.forwardRef((props, ref: React.Ref<HTMLDivElement>) => (
	<>
	</>
));

function NotFound() {
  return (
    <p>Not</p>
  );
}

function App() {
  const h1Ref = useRef<HTMLHeadingElement>(null)
   useLayoutEffect(() => {
	if (h1Ref.current) {
		h1Ref.current.style.backgroundColor = "black"
       		 h1Ref.current.style.width = "10%";
		h1Ref.current.style.borderRadius = "15px 50px 30px";
	}
    console.log(h1Ref); // { current: <h1_object> }
  })  
return (
    <div className="App">
	<Navi />	
    	 <BrowserRouter>
		<Routes>
			<Route path="/fortune" element={<Fortune />} />
		</Routes>
	</BrowserRouter>
	<footer></footer>	
    </div>
  );
}

export default App;
