<template>
  <div class="home-page">
    <div class="pro-sell">
      服务器商家：
      <el-select v-model="id" placeholder="请选择" @change="changeSell" :clearable="true">
        <el-option
          v-for="item in sellers"
          :key="item.ID"
          :label="item.SellerName"
          :value="item.ID">
        </el-option>
      </el-select>
    </div>

    <div class="pro-sell" style="margin: 0 10px">
      状态：
      <el-select v-model="status" placeholder="请选择" @change="changeSell" :clearable="true">
        <el-option
          v-for="item in statusFilter"
          :key="item.value"
          :label="item.name"
          :value="item.value">
        </el-option>
      </el-select>
      <el-button type="primary" @click="changeSell" :loading="loading" style="margin-left: 10px">{{
          btnName
        }}
      </el-button>
    </div>

    <el-skeleton
      :loading="loading"
      animated
      :throttle="500"
    >
      <template slot="template">
        <div style="padding: 14px;">
          <el-skeleton-item variant="h1"/>
          <el-skeleton-item variant="h1"/>
        </div>
      </template>
      <template>
        <el-table :data="tableData" style="width: 100%">
          <el-table-column prop="Name" label="名字" width="180"></el-table-column>
          <el-table-column prop="URL" label="链接" width="450"></el-table-column>
          <el-table-column prop="Price" label="价格" width="180"></el-table-column>
          <el-table-column prop="Stock" label="是否有货">
            <template slot-scope="scope">
              <el-tag
                :type="scope.row.Stock ? 'success' : 'danger'"
                disable-transitions>{{ scope.row.Stock ? '有货' : '缺货' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <template v-if="scope.row.Stock">
                <el-button @click="open_page(scope.row.URL)" type="success" size="small">购买</el-button>
              </template>
              <template v-else>
                <el-button type="danger" size="small" disabled>购买</el-button>
              </template>
            </template>
          </el-table-column>
        </el-table>
      </template>
    </el-skeleton>
  </div>
</template>
<script>
import { GetSellerApi, GetPlistApi } from '@/apis/data.api';
import useCounterStore from '@/store/modules/app';
import { useRoute } from '@/router';
import { open_page } from '@/utils/util';

export default {
  data() {
    return {
      loading: true,
      btnName: '加载中',
      testNum: 0,
      tableData: [],
      store: useCounterStore(),
      useRoute,
      open_page,
      sellers: [],
      statusFilter: [
        {
          name: '缺货',
          value: 1,
        },
        {
          name: '有货',
          value: 2,
        },
      ],
      id: undefined,
      status: undefined,
    };
  },
  mounted() {
    this.getSellers();
    this.getPlistApi();
  },
  computed: {},
  methods: {
    async getSellers() {
      let res = await GetSellerApi();
      this.sellers = res.data;
    },
    async getPlistApi() {
      this.loading = true;
      this.tableData = [];
      let res = await GetPlistApi({ id: this.id, stock: this.status });
      this.tableData = res.data;
      this.btnName = '刷新';
      this.loading = false;
    },

    changeSell() {
      this.getPlistApi();
    },
    handleClick(val) {
      open_page(val.URL);
    },
  },
  components: {},
};
</script>
<style lang="scss" scoped>
.home-page {
  .test-text {
    color: v-bind(textColor);
  }

  .pro-sell {
    display: inline-block;
    margin: 20px 0;
  }
}
</style>
