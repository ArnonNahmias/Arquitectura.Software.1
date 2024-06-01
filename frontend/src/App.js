// App.js
import React, { useState } from 'react';
import TopMenu from './components/TopMenu/TopMenu';
import CoursesAdmin from './components/Courses/CoursesAdmin';
import CommonUserView from './components/CommonUserView/CommonUserView';
import Foot from './components/Foot/Foot';
import UserValidation from './components/UserValidation/UserValidation';

function App() {
  const [userRole, setUserRole] = useState(null);

  const handleLogin = (role) => {
    setUserRole(role);
  };

  const handleSignOut = () => {
    setUserRole(null);
  };

  return (
    <div className="App">
      <TopMenu userRole={userRole} onSignOut={handleSignOut} />
      {userRole === null ? (
        <UserValidation onLogin={handleLogin} />
      ) : userRole === 'admin' ? (
        <CoursesAdmin />
      ) : (
        <CommonUserView />
      )}
      <Foot />
    </div>
  );
}

export default App;
