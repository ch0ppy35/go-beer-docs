import BeerList from './components/BeerList/BeerList';
import BreweryList from './components/BreweryList/BreweryList';
import Header from './components/Header/Header';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <Routes>
          <Route path="/" element={<BeerList />} />
          <Route path="/breweries" element={<BreweryList />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
