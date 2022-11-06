import React from 'react';

type BuildingProps = {
    onClick: React.MouseEventHandler<HTMLDivElement>;
    name: string;
    level: number;
};

const Building: React.FC<BuildingProps> = ({onClick, name, level}) => {
  return (
      <div onClick = {onClick}>
          {name} {level}
      </div>
  );
};

export default Building;