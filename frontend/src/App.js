import React from 'react';
import TopMenu from './components/TopMenu/TopMenu';
import Courses from './components/Courses/Courses';
import Foot from './components/Foot/Foot';


function App() {
  return (
    <div className="App">
      <TopMenu />
      <Courses />
      <Foot />
    </div>
  );
}

export default App;
