import { useState } from 'react';


import RoundButton from 'components/buttons/RoundButton';
import TitleCard from 'components/cards/TitleCard';
import InfoCard from 'components/cards/InfoCard';
import City from 'components/city/City';

import './App.scss';

function Show(id:string) {
  const x = document.getElementById(id);
  if(x != null){
    if (x.style.display === "none") {
      x.style.display = "block";
    }
  }
}

function Hide(id:string) {
  const x = document.getElementById(id);
  if(x != null){
    if (x.style.display === "block") {
      x.style.display = "none";
    }
  }
}

const App: React.FC<unknown> = () => {
  const [currentViewTitle, setCurrentViewTitle] = useState('City View');
  const [currentShow, setCurrentShow] = useState("display:none;");

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
        <TitleCard>{currentViewTitle}</TitleCard>
        <p id='test'>Test test, one two three...?</p>
        <div id = 'city' style = {{display: 'block'}}><City>City</City></div>
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
