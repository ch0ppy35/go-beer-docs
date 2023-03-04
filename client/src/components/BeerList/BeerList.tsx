import React, { useEffect, useState } from 'react';
import { BeersService, controllers_BeerResponse } from '../../generated';
import { Container, Table } from 'react-bootstrap';
import { BeatLoader } from 'react-spinners';


type BeerListProps = {};

const BeerList: React.FC<BeerListProps> = () => {
  const [beers, setBeers] = useState<Array<controllers_BeerResponse>>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const fetchBeers = async () => {
      try {
        const response = await BeersService.getBeers();
        setBeers(response);
      } catch (error: any) {
        if (error.response.status === 404) {
          setError('No Beers found!');
        } else {
          setError(error.message);
        }
      } finally {
        setIsLoading(false);
      }
    };

    fetchBeers();
  }, []);

  if (isLoading) {
    return (
      <div className="loader-container">
        <BeatLoader color={'#ffffff'} loading={isLoading} size={15} />
      </div>
    );
  }

  if (error) {
    return <p className="text-light">Failed to fetch beers: {error}</p>;
  }

  if (beers.length === 0) {
    return <p className="text-light">No beers found!</p>;
  }

  return (
    <Container className="my-4">
      <Table className="table-dark table-striped table-bordered table-hover">
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
      </Table>
    </Container>
  );
};

export default BeerList;
