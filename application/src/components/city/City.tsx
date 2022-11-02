import React from 'react';

import './City.scss';

type CityProps = {
  children: string;
  id: string;
};

const City: React.FC<CityProps> = ({children, id}) => {
  return (
    <button className='City' id = {id}>
      {children}
    </button>
  );
};

export default City;