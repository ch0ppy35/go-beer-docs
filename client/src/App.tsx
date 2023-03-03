import { useState, useEffect } from "react";
import "./App.css";
import { BeersService, controllers_BeerResponse } from "./generated"



const App = () => {

  const [beers, setBeers] = useState<Array<controllers_BeerResponse>>([]);

  useEffect(() => {
    BeersService.getBeers().then(response => {
      setBeers(response);
    });
  }, []);

  return (
    <table>
      <thead>
        <tr>
          <th>Beer Name</th>
          <th>Brewery Name</th>
        </tr>
      </thead>
      <tbody>
        {beers.map(beer => (
          <tr key={beer.name}>
            <td>{beer.name}</td>
            <td>{beer.brewery?.name}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default App;