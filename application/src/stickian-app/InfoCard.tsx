import React from 'react';

import './InfoCard.scss';

type InfoCardProps = {
  children: JSX.Element;
};

const InfoCard: React.FC<InfoCardProps> = ({ children }) => {
  return (
    <div className='InfoCard'>
      <p>{children}</p>
    </div>
  );
};

export default InfoCard;
