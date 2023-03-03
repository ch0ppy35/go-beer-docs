import React, { useEffect, useState } from 'react';
import { BeersService, controllers_BeerResponse } from '../../generated';

type BeerListProps = {};

const BeerList: React.FC<BeerListProps> = () => {
  const [beers, setBeers] = useState<Array<controllers_BeerResponse>>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    BeersService.getBeers()
      .then(response => {
        setBeers(response);
      })
      .catch(error => {
        setError(error.message);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  if (isLoading) {
    return <p>Loading beers...</p>;
  }

  if (error) {
    return <p>Failed to fetch beers: {error}</p>;
  }

  if (beers.length === 0) {
    return <p>No beers found!</p>;
  }

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
