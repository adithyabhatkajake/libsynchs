# import seaborn as sns
# data = sns.load_dataset("tips")
# filepath = '/root/.pip'
# fig_name = 'scatterplot.png'

# # fig_path为想要存入的文件夹或地址
# fig_path = filepath + '/' + fig_name
# fig = sns.scatterplot(x = data['total_bill'], y = data['tip'], hue = 'time', 
# data = data, palette = 'Set1', s = 100)
# scatter_fig = fig.get_figure()
# scatter_fig.savefig(fig_path, dpi = 400)

import seaborn as sns
import pandas as pd
import matplotlib.pyplot as plt
# Load an example datase
filepath = '/root/github.com/Reputation-based-state-machine-replication-/paperpicture'
fig_name = 'scatterplot2.png'
fig_path = filepath + '/' + fig_name
# Create a visualization
# 加载数据
flights=pd.read_csv('sa1.csv')
flights.head()
# 长型数据多折线图
dict1 = {'axes.axisbelow': True,  #轴在图形的下面
'axes.edgecolor': 'black',	#边框的颜色
'axes.facecolor': '#EAEAF2',	#背景颜色
'axes.grid': True,	#是否显示网格
'axes.labelcolor': '.15',
'axes.linewidth': 0.0,
'figure.facecolor': 'white',
'font.family': ['sans-serif'],
'font.sans-serif': ['Arial',
'Liberation Sans',
'Bitstream Vera Sans',
'sans-serif'],
'grid.color': 'white',
'grid.linestyle': '--',
'image.cmap': 'Greys',
'legend.frameon': False,
'legend.numpoints': 1,
'legend.scatterpoints': 1,
'lines.solid_capstyle': 'round',
'text.color': 'black',
'xtick.color': 'black',
'xtick.direction': 'out',
'xtick.major.size': 0.0,
'xtick.minor.size': 0.0,
'ytick.color': 'black',
'ytick.direction': 'out',
'ytick.major.size': 0.0,
'ytick.minor.size': 0.0}
sns.set(context='paper', 
palette='deep', font='sans-serif', font_scale=1, 
color_codes=True, rc=dict1)
sns.despine(offset=15)
sns.despine(fig=None, ax=None,
 top=True, right=True, left=True, bottom=True,
  offset=None, trim=False)
fig = sns.lineplot(data=flights,x='Throughput (Kops/sec)',y='Latency(ms)',dashes=False,sort=True,
errorbar=None,hue='BlockSize',style='BlockSize',markers=['^'],orient='y')
fig.set_xlim(10,310) 
fig.set_ylim(100,200)
plt.yscale('log')
lineplot_figure = fig.get_figure()
lineplot_figure.savefig(fig_path, dpi = 400)
plt.show()
