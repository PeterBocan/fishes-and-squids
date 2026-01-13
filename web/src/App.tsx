
import { useEffect, useState } from 'react';

class Product {
  id: string;
  name: string;
  species: string;
  originCountry: string;
  productionMethod: "aquaculture" | "wild-caught";
  pricePerTonne: number;
  supplierName: string;
  certification: string;

  constructor(
    id: string,
    name: string,
    species: string,
    originCountry: string,
    productionMethod: "aquaculture" | "wild-caught",
    pricePerTonne: number,
    supplierName: string,
    certification: string
  ) {
    this.id = id;
    this.name = name;
    this.species = species;
    this.originCountry = originCountry;
    this.productionMethod = productionMethod;
    this.pricePerTonne = pricePerTonne;
    this.supplierName = supplierName;
    this.certification = certification;
  }
}


function detectURLEndpoint(): string {
  if (window.location.host.startsWith("localhost")) {
    return "http://localhost:8080";
  }
  return window.location.host;
}

const API_URL = detectURLEndpoint();

function App() {
  console.log(API_URL);
  const [products, setProducts] = useState([] as Product[]);

  useEffect(() => {
      let isMounted = true; // prevents state updates after unmount

    async function fetchProducts() {
      try {
        const response = await fetch(API_URL + "/products");;

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();

        if (isMounted) {
          setProducts(data);
        }
      } catch (err) {
        console.debug(err);
      }
  }
  fetchProducts();

    return () => {
      isMounted = false;
    };

  }, []);

  return (
    <>
      <h1 className="text-3xl font-bold underline">Ocean Orchestra</h1>
         <ul>
        {products.map(p => (
          <li className="font-color: blue" key={p.id}> <span className="font-bold">{p.species}</span> / {p.name} </li>
        ))}
    </ul> 
    </>
  );
}

export default App;
