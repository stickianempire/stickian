import React from 'react';

import './Building.scss';

type BuildingProps = {
  onClick: React.MouseEventHandler<HTMLDivElement>;
  name: string;
  level: number;
  img: string;
};

const Building: React.FC<BuildingProps> = ({ onClick, name, level, img }) => {
  return (
    <div onClick={onClick} className='Building'>
      <div className='Cropper'>
        <img src={img} alt={'Img for ' + name + ' not found'} className='Cropped' />
      </div>
      {name} {level}
    </div>
  );
};

export default Building;
