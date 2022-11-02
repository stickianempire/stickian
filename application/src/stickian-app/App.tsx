import { useState } from 'react';


import RoundButton from 'components/buttons/RoundButton';
import TitleCard from 'components/cards/TitleCard';
import InfoCard from 'components/cards/InfoCard';
import City from 'components/city/City';

import './App.scss';

function Show(id:string) {
  const x = document.getElementById(id);
  if (x!.style.display === "none") {
    x!.style.display = "block";
  }
}

function Hide(id:string) {
  const x = document.getElementById(id);
  if (x!.style.display === "block") {
    x!.style.display = "none";
  }
}

const App: React.FC<unknown> = () => {
  const [curretViewTitle, setCurrentViewTitle] = useState('City View');

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
        <p id='test'>Resources</p>
        <InfoCard>
          <div>
            <p>Sticks / Stone / Clay</p>
          </div>
        </InfoCard>
      </div>
      <div className='MainView'>
        <TitleCard>{curretViewTitle}</TitleCard>
        <p id='test'>Test test, one two three...?</p>
        <City id='city'>City</City>
      </div>
      <div className='Footer'>
        <RoundButton id = 'button1' onClick={() => {setCurrentViewTitle('City View'); Show('city')}}>City</RoundButton>
        <RoundButton id = 'button2' onClick={() => {setCurrentViewTitle('Map View'); Hide('city')}}>Map</RoundButton>
        <RoundButton id = 'button3' onClick={() => {setCurrentViewTitle('Social'); Hide('city')}}>Social</RoundButton>
      </div>
    </div>
  );
};

export default App;
