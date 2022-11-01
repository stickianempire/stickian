import React from 'react';

import './RoundButton.scss';

type RoundButtonProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
  children: string;
};

const RoundButton: React.FC<RoundButtonProps> = ({ onClick, children }) => {
  return (
    <button className='RoundButton' onClick={onClick}>
      {children}
    </button>
  );
};

export default RoundButton;
