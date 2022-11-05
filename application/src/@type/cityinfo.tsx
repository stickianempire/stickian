type buildingInfo = {
  name: string;
  level: number;
  built: boolean;
  // etc.
};

type cityInfo = {
  cityHall: buildingInfo;
  barracks: buildingInfo;
  forge: buildingInfo;
  // etc.
};

export type { buildingInfo, cityInfo };