const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
module.exports = {
  entry: './src/js/app.js',
  mode: 'development',
  plugins: [
	new HtmlWebpackPlugin({
	  template: 'src/index.html'
	  }),
	  new MiniCssExtractPlugin()
	],
  module: {
	rules: [
	  {
		test: /\.css$/,
		use: [
			MiniCssExtractPlugin.loader,
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

