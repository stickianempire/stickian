import React from 'react';

import './RoundButton.scss';

type RoundButtonProps = {
  onClick: React.MouseEventHandler<HTMLButtonElement>;
  children: string;
  id: string;
};

const RoundButton: React.FC<RoundButtonProps> = ({ onClick, children, id }) => {
  return (
    <button className='RoundButton' onClick={onClick} id = {id}>
      {children}
    </button>
  );
};

export default RoundButton;