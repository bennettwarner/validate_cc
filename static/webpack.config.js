const HtmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {
  entry: './src/js/app.js',
  mode: 'development',
  plugins: [
	new HtmlWebpackPlugin({
	  template: 'src/index.html'
	  })
	],
  module: {
	rules: [
	  {
		test: /\.css$/,
		use: [
		  'style-loader',
		  'css-loader',
		],
	  },
	  {
		test: /\.(svg|gif|png|eot|woff|woff2|ttf)$/,
		use: [
		  'url-loader',
		],
	  },
	  
	],
  },
  output: {
	path: `${__dirname}/dist`,
	filename: 'bundle.js',
  },
};

