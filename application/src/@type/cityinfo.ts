type buildingInfo = {
  name: string;
  level: number;
  available: boolean;
  // etc.
};

type cityInfo = {
  cityHall: buildingInfo;
  barracks: buildingInfo;
  forge: buildingInfo;
  wall: buildingInfo; 
  // etc.
};

export type { buildingInfo, cityInfo };
