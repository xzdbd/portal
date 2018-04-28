import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import Typography from 'material-ui/Typography'
import Grid from 'material-ui/Grid'
import SearchBar from 'material-ui-search-bar'

const styles = {
  root: {
    flexGrow: 1,
    paddingTop: 50,
    background: 'url(http://www.sambusgeospatial.com/wp-content/uploads/2015/04/online-banner.jpg)',
    minHeight: 380
  },
  title: {
    margin: 30
  },
  container: {
    width: '100%'
  }
}

function HeaderBar (props) {
  const { classes } = props
  return (
    <div className={classes.root}>
      <Grid container spacing={40} className={classes.container} justify='center'>
        <Grid item xs={12} className={classes.title}>
          <Typography variant='display1' color='inherit' align='center'>
            <b>Get Sharing Files with Portal</b>
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <SearchBar
            placeholder='请输入订单id'
            onChange={() => console.log('onChange')}
            onRequestSearch={() => console.log('onRequestSearch')}
            style={{
              margin: '0 auto',
              maxWidth: 800
            }}
          />
        </Grid>
        {/* <Grid item xs={8}>
                <img src="https://slidesup.com/static/media/hero.4744b570.svg" width="100%" height="200px"/>
            </Grid> */}
      </Grid>
    </div>
  )
}

HeaderBar.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(HeaderBar)
