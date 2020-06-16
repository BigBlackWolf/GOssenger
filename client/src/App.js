import React from 'react';
import './App.css';
import {Container} from "semantic-ui-react";
import TODOList from "./TODOList";

function App() {
  return (
    <div>
      <Container>
        <TODOList />
      </Container>
    </div>
  );
}

export default App;
