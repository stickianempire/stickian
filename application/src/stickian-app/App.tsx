import { useState } from 'react';

import RoundButton from 'components/buttons/RoundButton';
import TitleCard from 'stickian-app/TitleCard';
import InfoCard from 'stickian-app/InfoCard';

import './App.scss';

const App: React.FC<unknown> = () => {
  const [curretViewTitle, setCurrentViewTitle] = useState('Test test, one two three...?');

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
      <div className='MainView'>
        <TitleCard>{curretViewTitle}</TitleCard>
        <p id='test'>Test test, one two three...?</p>
      </div>
      <div className='Footer'>
        <RoundButton onClick={() => setCurrentViewTitle('City View')}>City</RoundButton>
        <RoundButton onClick={() => setCurrentViewTitle('Map View')}>Map</RoundButton>
        <RoundButton onClick={() => setCurrentViewTitle('Social')}>Social</RoundButton>
      </div>
    </div>
  );
};

export default App;
