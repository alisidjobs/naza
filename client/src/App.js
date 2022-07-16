import './App.css'
import React from 'react'
import {Container, Checkbox , Button} from "semantic-ui-react";
import FormPerso from "./FormPerso";

function App() {
  return (
    <div className="App">
        <div className="test">
            <Container>
                <FormPerso />
            </Container>
        </div>
    </div>
  );
}

export default App;
