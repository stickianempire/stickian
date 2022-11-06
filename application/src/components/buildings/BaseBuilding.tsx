import React from 'react';

import './BaseBuilding.scss';

type BuildingProps = {
  onClick: React.MouseEventHandler<HTMLDivElement>;
  name: string;
  level: number;
  img: string;
  top: number;
  left: number;
};

const Building: React.FC<BuildingProps> = ({ onClick, name, level, img, top, left }) => {
  return (
    <div onClick={onClick} className='BaseBuilding'>
      <div
        className='Cropper'
        style={{ background: 'url("' + img + '") ' + top + 'px ' + left + 'px' }}
      ></div>
      {name} {level}
    </div>
  );
};

export default Building;
