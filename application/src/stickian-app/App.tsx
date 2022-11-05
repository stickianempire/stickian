import { useState } from 'react';

import RoundButton from 'components/buttons/RoundButton';
import InfoCard from 'components/cards/InfoCard';
import CityView from 'components/views/CityView';
import MapView from 'components/views/MapView';
import SocialView from 'components/views/SocialView';

import './App.scss';

const enum viewID {
  City,
  Map,
  Social,
}

const App: React.FC<unknown> = () => {
  const [currentView, setCurrentView] = useState(viewID.City);

  const _displaySelectedView = (selected: viewID) => {
    switch (selected) {
      case viewID.City:
        return <CityView />;
      case viewID.Map:
        return <MapView />;
      case viewID.Social:
        return <SocialView />;
    }
  };

  return (
    <div className='App'>
      <div className='LeftPad'>
        <p id='test'>This stuff is on the left</p>
        <InfoCard>
          <p>This is an info card that has a very long, long, long description.</p>
        </InfoCard>
        <InfoCard>
          <div>
            <p>This is another info card with one section here</p>
            <p>And another sectiont here</p>
          </div>
        </InfoCard>
      </div>
      <div className='RightPad'>
        <p id='test'>This stuff is on the right</p>
      </div>
      <div className='MainView'>{_displaySelectedView(currentView)}</div>
      <div className='Footer'>
        <RoundButton onClick={() => setCurrentView(viewID.City)}>City</RoundButton>
        <RoundButton onClick={() => setCurrentView(viewID.Map)}>Map</RoundButton>
        <RoundButton onClick={() => setCurrentView(viewID.Social)}>Social</RoundButton>
      </div>
    </div>
  );
};

export default App;
