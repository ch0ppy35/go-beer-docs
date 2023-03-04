import { BreweriesService, controllers_BreweryResponse } from '../../generated';
import React, { useEffect, useState } from 'react';
import { Container, Table } from 'react-bootstrap';

type BreweryListProps = {};

const BreweryList: React.FC<BreweryListProps> = () => {
  const [breweries, setBreweries] = useState<Array<controllers_BreweryResponse>>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const fetchBreweries = async () => {
      try {
        const response = await BreweriesService.getBreweries();
        setBreweries(response);
      } catch (error: any) {
        if (error.response.status === 404) {
          setError('No breweries found!');
        } else {
          setError(error.message);
        }
      } finally {
        setIsLoading(false);
      }
    };

    fetchBreweries();
  }, []);

  if (isLoading) {
    return <p className="text-light">Loading breweries...</p>;
  }

  if (error) {
    return <p className="text-light">Failed to fetch breweries: {error}</p>;
  }

  if (breweries.length === 0) {
    return <p className="text-light">No breweries found!</p>;
  }

  return (
    <Container className="my-4">
      <Table className="table-dark table-striped table-bordered table-hover">
        <thead>
          <tr>
            <th>Brewery Name</th>
          </tr>
        </thead>
        <tbody>
          {breweries.map(brewery => (
            <tr key={brewery.name}>
              <td>{brewery.name}</td>
            </tr>
          ))}
        </tbody>
      </Table>
    </Container>
  );
};

export default BreweryList;
