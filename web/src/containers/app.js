import React from 'react'
import PropTypes from 'prop-types'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import Grid from 'material-ui/Grid'
import HeaderBar from '../components/header'
import Search from '../components/search'
import ItemList from '../components/itemList'
import Footer from '../components/footer'
import * as ItemActions from '../actions'

const App = (props) => (
  <div>
    <Grid container spacing={0} justify='center'>
      <Grid item xs={12} >
        <HeaderBar />
      </Grid>
      <Grid item xs={12}>
        <Search />
      </Grid>
      <Grid item xs={12}>
        <ItemList />
      </Grid>
      <Grid item xs={12}>
        <Footer />
      </Grid>
    </Grid>
  </div>
)

App.propTypes = {
  items: PropTypes.array.isRequired,
  actions: PropTypes.object.isRequired
}

const mapStateToProps = state => ({
  items: state.items
})

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ItemActions, dispatch)
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
