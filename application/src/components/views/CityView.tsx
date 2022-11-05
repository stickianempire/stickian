import React, { useState, useEffect } from 'react';

import { cityInfo, buildingInfo } from '@type/cityinfo';

import TitleCard from 'components/cards/TitleCard';

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
      forge: { name: 'Forge', level: 0, built: false },
    });
  }, []);

  const _renderBuildingIfBuilt = (currentBuildingInfo: buildingInfo) => {
    if (currentBuildingInfo.built) {
      return (
        <p>
          I got a {currentBuildingInfo.name} at level {currentBuildingInfo.level}
        </p>
      );
    }
    return;
  };

  return (
    <div>
      <TitleCard>City view </TitleCard>
      {_renderBuildingIfBuilt(currentCityInfo.cityHall)}
      {_renderBuildingIfBuilt(currentCityInfo.barracks)}
      {_renderBuildingIfBuilt(currentCityInfo.forge)}
    </div>
  );
};

export default CityView;
