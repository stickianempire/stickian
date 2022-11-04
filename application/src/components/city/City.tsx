import React from 'react';

import './City.scss';

type CityProps = {
  children: string;
};

const City: React.FC<CityProps> = ({children}) => {
  return (
    <button className='City'>
      {children}
    </button>
  );
};

export default City;