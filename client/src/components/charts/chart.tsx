import * as React from 'react';

import * as Plotly from 'plotly.js';

import PlotlyChart from './plotlyChart';

interface Props {
  data: Plotly.PlotData[];
  title: string;
}

export default class Chart extends React.Component<Props, {}> {
  public render() {
    return (
      <PlotlyChart
        data={this.props.data}
        layout={{title: this.props.title, autosize: true, titlefont: {size: 13}}}
        config={{displayModeBar: false}}
      />
    );
  }
}
