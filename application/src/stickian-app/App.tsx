import { useEffect, useState } from 'react';

import RoundButton from 'components/buttons/RoundButton';
import TitleCard from 'components/cards/TitleCard';
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
  const [count, setCount] = useState(0);

  useEffect(() => {
    const timer = window.setInterval(() => {
      setCount((count) => count + 1);
    }, 1000);
    return () => {
      window.clearInterval(timer);
    };
  }, []);

  const _displaySelectedView = (selected: viewID) => {
    switch (selected) {
      case viewID.City:
        return (
          <div className='MainView'>
            <TitleCard>City View</TitleCard>
            <CityView />
          </div>
        );
      case viewID.Map:
        return (
          <div className='MainView'>
            <TitleCard>Map View</TitleCard>
            <MapView />
          </div>
        );
      case viewID.Social:
        return (
          <div className='MainView'>
            <TitleCard>Social</TitleCard>
            <SocialView />
          </div>
        );
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
        <p id='test'>This stuff is on the right {count}</p>
      </div>
      {_displaySelectedView(currentView)}
      <div className='Footer'>
        <RoundButton onClick={() => setCurrentView(viewID.City)}>City</RoundButton>
        <RoundButton onClick={() => setCurrentView(viewID.Map)}>Map</RoundButton>
        <RoundButton onClick={() => setCurrentView(viewID.Social)}>Social</RoundButton>
      </div>
    </div>
  );
};

export default App;
