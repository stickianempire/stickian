import React, { useState, useEffect } from 'react';

import { cityInfo, buildingInfo } from '@type/cityinfo';

import TitleCard from 'components/cards/TitleCard';
import Building from 'components/buildings/Building';

import './CityView.scss';

type CityViewProps = unknown;

const CityView: React.FC<CityViewProps> = () => {
  // TODO set current city should depend on callbacks from API
  const [currentCityInfo, setCurrentCityInfo] = useState<cityInfo>({
    cityHall: { name: 'City Hall', level: 0, available: false },
    barracks: { name: 'Barracks', level: 0, available: false },
    forge: { name: 'Forge', level: 0, available: false },
    wall: { name: 'Wall', level: 0, available: false },
  });

  useEffect(() => {
    // This would be an API call when mounting component
    setCurrentCityInfo({
      cityHall: { name: 'City Hall', level: 5, available: true },
      barracks: { name: 'Barracks', level: 2, available: true },
      forge: { name: 'Forge', level: 0, available: false },
      wall: { name: 'Wall', level: 0, available: false },
    });
  }, []);

  const _renderBuildingIfavailable = (currentBuildingInfo: buildingInfo, buildingImg: string) => {
    if (currentBuildingInfo.available) {
      return (
        <Building
          onClick={() => {
            return;
          }}
          name={currentBuildingInfo.name}
          level={currentBuildingInfo.level}
          img={buildingImg}
        ></Building>
      );
    }
    return;
  };

  return (
    <div>
      <TitleCard>City view</TitleCard>
      {_renderBuildingIfavailable(currentCityInfo.cityHall, '/images/city.png')}
      {_renderBuildingIfavailable(currentCityInfo.barracks, '/images/city.png')}
      {_renderBuildingIfavailable(currentCityInfo.forge, '/images/city.png')}
    </div>
  );
};

export default CityView;
