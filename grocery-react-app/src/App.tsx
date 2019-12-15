import React from 'react';
import logo from './logo.svg';
import './App.css';
import HeaderTab from './components/HeaderTab';
import SearchContainer from './components/SearchContainer';

const App: React.FC = () => {
    return (
        <div>
            <HeaderTab />
            <SearchContainer />
        </div>
    );
}

export default App;
