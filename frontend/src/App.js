import React from 'react';
import TopMenu from './components/TopMenu/TopMenu';
import CoursesAdmin from './components/Courses/CoursesAdmin';
import Foot from './components/Foot/Foot';
import UserValidation from './components/UserValidation/UserValidation';


function App() {
  return (
    <div className="App">
      <TopMenu />
      <UserValidation />
      <CoursesAdmin />
      <Foot />
    </div>
  );
}

export default App;
