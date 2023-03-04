import BeerList from './components/BeerList/BeerList';
import BreweryList from './components/BreweryList/BreweryList';
import Header from './components/Header/Header';
import 'bootstrap/dist/css/bootstrap.min.css';


function App() {
  return (
    <div className="App">
      <Header />
      <BeerList />
      <BreweryList />
    </div>
  );
}

export default App;
