import React, { useState, useEffect } from 'react';

import { cityInfo, buildingInfo } from '@type/cityinfo';

import TitleCard from 'components/cards/TitleCard';

import './CityView.scss';
import Building from 'components/buildings/Building';

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
      forge: { name: 'Forge', level: 0, available: false},
      wall: { name: 'Wall', level: 0, available: false},
    });
  }, []);

  const _renderBuildingIfavailable = (currentBuildingInfo: buildingInfo, buildingId:string) => {
    if (currentBuildingInfo.available) {
      return (
        <div id = {buildingId}>
          <Building onClick={() => {return;}} name = {currentBuildingInfo.name} level = {currentBuildingInfo.level}></Building>
        </div>
      );
    }
    return;
  };

  return (
    <div>
      <TitleCard>City view</TitleCard>
      {_renderBuildingIfavailable(currentCityInfo.cityHall, 'City-Hall')}
      {_renderBuildingIfavailable(currentCityInfo.barracks, 'Barracks')}
      {_renderBuildingIfavailable(currentCityInfo.forge, 'Forge')}
    </div>
  );
};

export default CityView;
