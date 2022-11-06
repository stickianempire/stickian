import React, { useState, useEffect } from 'react';

import { cityInfo, buildingInfo } from '@type/cityinfo';

import TitleCard from 'components/cards/TitleCard';
import Building from 'components/buildings/BaseBuilding';

import './CityView.scss';

type CityViewProps = unknown;

const CityView: React.FC<CityViewProps> = () => {
  // TODO set current city should depend on callbacks from API
  const [currentCityInfo, setCurrentCityInfo] = useState<cityInfo>({
    cityHall: { level: 0, available: false },
    barracks: { level: 0, available: false },
    forge: { level: 0, available: false },
    wall: { level: 0, available: false },
  });

  useEffect(() => {
    // This would be an API call when mounting component
    setCurrentCityInfo({
      cityHall: { level: 5, available: true },
      barracks: { level: 2, available: true },
      forge: { level: 0, available: false },
      wall: { level: 0, available: false },
    });
  }, []);

  const _renderBaseBuildingIfAvailable = (
    name: string,
    currentBuildingInfo: buildingInfo,
    buildingImg: string,
    top: number,
    left: number,
  ) => {
    if (currentBuildingInfo.available) {
      return (
        <Building
          onClick={() => {
            console.log('you clicked ' + name);
            return;
          }}
          name={name}
          level={currentBuildingInfo.level}
          img={buildingImg}
          top={top}
          left={left}
        ></Building>
      );
    }
    return;
  };

  return (
    <div>
      <TitleCard>City view</TitleCard>
      {_renderBaseBuildingIfAvailable(
        'City Hall',
        currentCityInfo.cityHall,
        '/images/city.png',
        0,
        0,
      )}
      {_renderBaseBuildingIfAvailable(
        'Barracks',
        currentCityInfo.barracks,
        '/images/city.png',
        -40,
        0,
      )}
      {_renderBaseBuildingIfAvailable(
        'The Forge',
        currentCityInfo.forge,
        '/images/city.png',
        100,
        200,
      )}
    </div>
  );
};

export default CityView;
