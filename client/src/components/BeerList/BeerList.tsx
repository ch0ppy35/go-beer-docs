import React, { useEffect, useState } from 'react';
import { BeersService, controllers_BeerResponse } from '../../generated';

type BeerListProps = {};

const BeerList: React.FC<BeerListProps> = () => {
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

export default BeerList;
