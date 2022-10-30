import React from 'react';

import './TitleCard.scss';

type TitleCardProps = {
  children: string;
};

const TitleCard: React.FC<TitleCardProps> = ({ children }) => {
  return (
    <div className='TitleCard'>
      <h1>{children}</h1>
    </div>
  );
};

export default TitleCard;
