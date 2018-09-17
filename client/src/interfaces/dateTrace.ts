import { ChartModes, ChartTypes } from '../models/chart';

export interface DateTrace {
  x: Plotly.Datum[];
  y: Plotly.Datum[];
  z?: Plotly.Datum[];
}

export interface Trace extends DateTrace {
  name: string;
  mode: ChartModes;
  type: ChartTypes;
  line: Partial<Plotly.ScatterLine>;
  // 'line.shape': string;
  // 'line.width': number;
  // 'line.smoothing': number;
  // 'line.color': Plotly.Color;
}
