import React, { useState, useEffect } from 'react';

import { cityInfo, buildingInfo } from '@type/cityinfo';

import TitleCard from 'components/cards/TitleCard';

import './CityView.scss';

type CityViewProps = unknown;

const CityView: React.FC<CityViewProps> = () => {
  // TODO set current city should depend on callbacks from API
  const [currentCityInfo, setCurrentCityInfo] = useState<cityInfo>({
    cityHall: { name: 'City Hall', level: 0, built: false },
    barracks: { name: 'Barracks', level: 0, built: false },
    forge: { name: 'Forge', level: 0, built: false },
  });

  useEffect(() => {
    // This would be an API call when mounting component
    setCurrentCityInfo({
      cityHall: { name: 'City Hall', level: 10, built: true },
      barracks: { name: 'Barracks', level: 2, built: true },
      forge: { name: 'Forge', level: 0, built: false},
    });
  }, []);

  const _renderBuildingIfBuilt = (currentBuildingInfo: buildingInfo, buildingId:string) => {
    if (currentBuildingInfo.built) {
      return (
        <div id = {buildingId}>
          <p>
            I got a {currentBuildingInfo.name} at level {currentBuildingInfo.level}
          </p>
        </div>
      );
    }
    return;
  };

  /*const _renderCity = () => {
    return(
      <img src = {city} alt = 'City' width='748px' height='537px'></img>
    );
  }*/

  const _renderBuildingImage = () => {
    return(
      <div>
        <div id="cityhallimg"></div>
        <div id="barracksimg"></div>
        <div id="forgeimg"></div>
      </div>
    );
  }

  return (
    <div>
      <TitleCard>City view</TitleCard>
      {_renderBuildingIfBuilt(currentCityInfo.cityHall, 'cityhall')}
      {_renderBuildingIfBuilt(currentCityInfo.barracks, 'barracks')}
      {_renderBuildingIfBuilt(currentCityInfo.forge, 'forge')}
      {_renderBuildingImage()}
    </div>
  );
};

export default CityView;
